syntax = "proto3";

package {{$.PackageName}}query;

option go_package = "pkg/{{$.PackageName}}query";

service {{$.StreamName}}Service {
    rpc Find(FindRequest) returns (FindResponse);
    rpc FindOne(FindOneRequest) returns (FindOneResponse);
}

message Instance {
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
   Instance result = 1;
}