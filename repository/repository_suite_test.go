package repository_test

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type ContainerAddress struct {
	Host      string
	Port      string
	Terminate func()
}

var (
	Redis ContainerAddress
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var _ = BeforeSuite(func() {
	logrus.Info("BeforeSuite ")
})

func setupRedisContainer() ContainerAddress {
	ctx := context.Background()
	redisContainerRequest := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections").WithStartupTimeout(time.Second * 300),
	}

	redisContainer, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: redisContainerRequest,
		Started:          true,
	})

	if err != nil {
		logrus.Fatalf("error starting redis container: %s", err)
	}

	redisHost, _ := redisContainer.Host(ctx)
	redisPort, err := redisContainer.MappedPort(ctx, "6379")
	if err != nil {
		logrus.Fatalf("redisContainer.MappedPort: %s", err)
	}

	terminateContainer := func() {
		logrus.Info("terminating es container...")
		if err := redisContainer.Terminate(ctx); err != nil {
			logrus.Fatalf("error terminating es container: %v\n", err)
		}
	}

	return ContainerAddress{redisHost, redisPort.Port(), terminateContainer}
}
