package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	subscription := "CID_weilin_test_996"
	//clusterName := "DefaultCluster"
	//nameSrvAddr := []string{"127.0.0.1:9876"}

	testAdmin, err := admin.NewAdmin(
		//admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)),
		admin.WithResolver(primitive.NewHttpResolver("DefaultNameserver")),
		admin.WithCredentials(primitive.Credentials{
			AccessKey: "RocketMQ",
			SecretKey: "12345678",
		}),
	)

	//create subscription
	err = testAdmin.CreateSubscription(
		context.Background(),
		admin.WithSubscription(subscription),
		admin.WithSubscriptionClusterName("taobaodaily-e2e-test"),
	)
	if err != nil {
		fmt.Println("Create subscription error:", err.Error())
	}
	err = testAdmin.DeleteSubscription(
		context.Background(),
		admin.WithSubscription(subscription),
		admin.WithSubscriptionClusterName("taobaodaily-e2e-test"),
		admin.WithSubscriptionCleanOffset(true),
	)
	if err != nil {
		fmt.Println("Delete subscription error:", err.Error())
	}
	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}
