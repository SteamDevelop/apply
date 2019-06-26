/**
 * @fileoverview gRPC-Web generated client stub for hidot
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.hidot = require('./hi_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.hidot.HiDotClient =
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
proto.hidot.HiDotPromiseClient =
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
 *   !proto.hidot.ReqData,
 *   !proto.hidot.ResData>}
 */
const methodInfo_HiDot_Hi = new grpc.web.AbstractClientBase.MethodInfo(
  proto.hidot.ResData,
  /** @param {!proto.hidot.ReqData} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.hidot.ResData.deserializeBinary
);


/**
 * @param {!proto.hidot.ReqData} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.hidot.ResData)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.hidot.ResData>|undefined}
 *     The XHR Node Readable Stream
 */
proto.hidot.HiDotClient.prototype.hi =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/hidot.HiDot/Hi',
      request,
      metadata || {},
      methodInfo_HiDot_Hi,
      callback);
};


/**
 * @param {!proto.hidot.ReqData} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.hidot.ResData>}
 *     A native promise that resolves to the response
 */
proto.hidot.HiDotPromiseClient.prototype.hi =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/hidot.HiDot/Hi',
      request,
      metadata || {},
      methodInfo_HiDot_Hi);
};


module.exports = proto.hidot;

