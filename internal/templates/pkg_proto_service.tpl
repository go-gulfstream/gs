syntax = "proto3";

package proto;

option go_package = "pkg/proto";

service {{$.StreamName}}Service {
    rpc Find(FindRequest) returns (FindResponse);
    rpc FindOne(FindOneRequest) returns (FindOneResponse);
}

message {{$.StreamName}} {
}

message FindRequest {
}

message FindResponse {
}

message FindOneRequest {
   string projection_id = 1;
   int64 version = 2;
}

message FindOneResponse {
   {{$.StreamName}} result = 1;
}