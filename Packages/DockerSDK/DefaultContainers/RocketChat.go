package DefaultContainers

import (
	constants "thrust/Packages/Constants"
	"thrust/Packages/DockerSDK"
	models "thrust/Packages/Models"
	"thrust/Utils"

	"github.com/docker/go-connections/nat"
)

func LaunchRocketChatContainer(sdk DockerSDK.Docker, networkID string) (string, error) {

	rocketChatContainer := models.Container{
		NetworkID: networkID, ContainerName: "rocketchat_" + Utils.RandomString(5), Image: constants.RocketChatImage, ExposedPort: nat.PortSet{
			"3000/tcp": {},
		}, PortBindings: nat.PortMap{
			"3000/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "3000",
				},
			},
		}, Env: []string{
			"MONGO_URL=mongodb://mongo:27017/rocketchat?replicaSet=rs0",
			"MONGO_OPLOG_URL=mongodb://mongo:27017/local?replicaSet=rs0",
			"ROOT_URL=http://localhost:3000",
			"PORT=3000",
			"DEPLOY_METHOD=docker",
			"OVERWRITE_SETTING_Apps_Framework_Development_Mode=true",
			"OVERWRITE_SETTING_Show_Setup_Wizard=Completed",
		}, Volumes: nil, Binds: nil, Aliases: []string{
			"rocketchat",
		}, Links: nil, Mount: nil, Commands: nil, Stdout: false,
	}

	containerID, err := sdk.CreateContainer(rocketChatContainer, false)

	return containerID, err
}
