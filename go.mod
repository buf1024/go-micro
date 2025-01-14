module github.com/buf1024/go-micro

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/beevik/ntp v0.2.0
	github.com/bitly/go-simplejson v0.5.0
	github.com/bwmarrin/discordgo v0.19.0
	github.com/cloudflare/cloudflare-go v0.10.2
	github.com/coreos/etcd v3.3.15+incompatible
	github.com/forestgiant/sliceutil v0.0.0-20160425183142-94783f95db6c
	github.com/fsnotify/fsnotify v1.4.7
	github.com/fsouza/go-dockerclient v1.4.4
	github.com/ghodss/yaml v1.0.0
	github.com/go-acme/lego/v3 v3.1.0
	github.com/go-log/log v0.1.0
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/websocket v1.4.1
	github.com/hashicorp/consul/api v1.2.0
	github.com/hashicorp/hcl v1.0.0
	github.com/imdario/mergo v0.3.7
	github.com/joncalhoun/qson v0.0.0-20170526102502-8a9cab3a62b1
	github.com/json-iterator/go v1.1.7
	github.com/lucas-clemente/quic-go v0.12.0
	github.com/mholt/certmagic v0.8.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.10.1-0.20190914150736-364c5a486180
	github.com/micro/mdns v0.3.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/nats-io/nats.go v1.8.1
	github.com/nlopes/slack v0.6.0
	github.com/pkg/errors v0.8.1
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7
	golang.org/x/net v0.0.0-20190930134127-c5a3c61f89f3
	google.golang.org/grpc v1.23.1
	gopkg.in/go-playground/validator.v9 v9.29.1
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/telegram-bot-api.v4 v4.6.4
)

replace github.com/micro/go-micro => ../go-micro
