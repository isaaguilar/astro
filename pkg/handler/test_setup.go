package handler

import (
	"os"

	"github.com/golang/mock/gomock"
	"github.com/reactiveops/dd-manager/pkg/datadog"
	"github.com/reactiveops/dd-manager/pkg/kube"
	mocks "github.com/reactiveops/dd-manager/pkg/mocks"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func setupTests(ctrl *gomock.Controller) (*kube.ClientInstance, *mocks.MockClientAPI) {
	os.Setenv("DEFINITIONS_PATH", "../config/test_conf.yml")
	os.Setenv("DD_API_KEY", "test")
	os.Setenv("DD_APP_KEY", "test")

	kubeClient := kube.ClientInstance{
		Client: fake.NewSimpleClientset(),
	}
	kube.SetInstance(kubeClient)

	ddMon := datadog.GetInstance()
	ddMock := mocks.NewMockClientAPI(ctrl)
	ddMon.Datadog = ddMock

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "foo",
		},
	}
	kubeClient.Client.CoreV1().Namespaces().Create(ns)

	return &kubeClient, ddMock
}
