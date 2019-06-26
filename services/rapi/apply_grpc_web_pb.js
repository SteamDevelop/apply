/**
 * @fileoverview gRPC-Web generated client stub for rapi
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.rapi = require('./apply_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rapi.RapiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rapi.RapiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rapi.ApplyReq,
 *   !proto.rapi.ApplyRes>}
 */
const methodInfo_Rapi_apply = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rapi.ApplyRes,
  /** @param {!proto.rapi.ApplyReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rapi.ApplyRes.deserializeBinary
);


/**
 * @param {!proto.rapi.ApplyReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rapi.ApplyRes)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rapi.ApplyRes>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rapi.RapiClient.prototype.apply =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rapi.Rapi/apply',
      request,
      metadata || {},
      methodInfo_Rapi_apply,
      callback);
};


/**
 * @param {!proto.rapi.ApplyReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rapi.ApplyRes>}
 *     A native promise that resolves to the response
 */
proto.rapi.RapiPromiseClient.prototype.apply =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rapi.Rapi/apply',
      request,
      metadata || {},
      methodInfo_Rapi_apply);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rapi.SignReq,
 *   !proto.rapi.SignRes>}
 */
const methodInfo_Rapi_sign = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rapi.SignRes,
  /** @param {!proto.rapi.SignReq} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rapi.SignRes.deserializeBinary
);


/**
 * @param {!proto.rapi.SignReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rapi.SignRes)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rapi.SignRes>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rapi.RapiClient.prototype.sign =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rapi.Rapi/sign',
      request,
      metadata || {},
      methodInfo_Rapi_sign,
      callback);
};


/**
 * @param {!proto.rapi.SignReq} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rapi.SignRes>}
 *     A native promise that resolves to the response
 */
proto.rapi.RapiPromiseClient.prototype.sign =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rapi.Rapi/sign',
      request,
      metadata || {},
      methodInfo_Rapi_sign);
};


module.exports = proto.rapi;

