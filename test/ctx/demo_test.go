package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/smallnest/rpcx/client"
)

func Test_Ctx(t *testing.T) {

	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Hour)

	deadline, ok := ctx.Deadline()
	fmt.Println("ok-->", ok)
	fmt.Println("deadline-->", deadline)

	duration := time.Until(deadline)

	if duration < 2*time.Hour {
		fmt.Println("可以的")
	}

	fmt.Println("duration-->", duration)

	xClient := &client.Client{}

	err := xClient.Call(ctx, "", "UserEvent", nil, nil)
	if err != nil {
		panic(err)
	}

}
