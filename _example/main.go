package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	v1 "github.com/erikh/k8s-api/api/v1"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	// don't take your crypto advice from me
	rand.Seed(time.Now().Unix())
}

func main() {
	k8s, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		panic(err)
	}

	u := uuid.New()

	ur := &v1.UUID{
		Spec: v1.UUIDSpec{
			UUID:      u.String(),
			RandomInt: rand.Int(),
		},
	}

	ur.SetNamespace(os.Getenv("KUBE_NAMESPACE"))
	ur.SetName(u.String())

	s := runtime.NewScheme()
	v1.AddToScheme(s)
	k8sClient, err := client.New(k8s, client.Options{Scheme: s})
	if err != nil {
		panic(err)
	}

	if err := k8sClient.Create(context.Background(), ur); err != nil {
		panic(err)
	}
}
