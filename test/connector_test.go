package main

import (
	"github.com/iamwm/connector"
	"testing"
)

func TestMongoConnect(t *testing.T) {
	type args struct {
		confPath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"mongo connect success", args{""}, true},
		{"mongo connect success with relative path", args{"./conf/mongo.yaml"}, true},
		{"mongo connect fail with wrong path", args{"./"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connector.MongoConnect(tt.args.confPath); got != tt.want {
				t.Errorf("MongoConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRabbitmqConnect(t *testing.T) {
	type args struct {
		confPath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"rabbitmq connect success", args{""}, true},
		{"rabbitmq connect success with relative path", args{"./conf/rabbitmq.yaml"}, true},
		{"rabbitmq connect fail with wrong path", args{"./"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connector.RabbitmqConnect(tt.args.confPath); got != tt.want {
				t.Errorf("RabbitmqConnect() = %v, want %v", got, tt.want)
			}
		})
	}
}
