/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  HelloRequest,
  HelloResponse} from './hello_service_pb';

export class HelloClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSay = new grpcWeb.AbstractClientBase.MethodInfo(
    HelloResponse,
    (request: HelloRequest) => {
      return request.serializeBinary();
    },
    HelloResponse.deserializeBinary
  );

  say(
    request: HelloRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: HelloResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/proto.Hello/Say',
      request,
      metadata || {},
      this.methodInfoSay,
      callback);
  }

}

