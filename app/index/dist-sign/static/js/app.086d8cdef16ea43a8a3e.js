webpackJsonp([1],{NHnr:function(e,r,t){"use strict";Object.defineProperty(r,"__esModule",{value:!0});var i=t("7+uW"),n={render:function(){var e=this.$createElement,r=this._self._c||e;return r("div",{attrs:{id:"app"}},[r("router-view")],1)},staticRenderFns:[]};var o=t("VU/8")({name:"App"},n,!1,function(e){t("gsu9")},null,null).exports,a=t("/ocq"),p=t("Zrlr"),s=t.n(p),l=t("wxAW"),u=t.n(l),g=t("OJMy"),c=t("gYUs"),d=function(){function e(){s()(this,e)}return u()(e,null,[{key:"apply",value:function(e){var r=new c.ApplyReq,t=new g.RapiPromiseClient("https://apply.scry.info:446/rapi",null,null);return r.setPass(e.password),r.setFinger(e.finger),t.apply(r,null)}},{key:"sign",value:function(e){var r=new c.SignReq,t=new g.RapiPromiseClient("https://apply.scry.info:446/rapi",null,null);return r.setAddr(e.addr),r.setFinger(e.finger),t.sign(r,null)}}]),e}(),y={name:"HelloWorld",data:function(){return{password:"",showControl:!0}},methods:{sign:function(){null===localStorage.getItem("sign")&&d.sign({password:this.password,finger:"sign"}).then(function(e){localStorage.setItem("sign","signed")}).catch(function(e){alert("操作失败！",e)})}},created:function(){null===localStorage.getItem("apply")&&(alert("你还没有报名！"),this.showControl=!1),null!==localStorage.getItem("sign")&&(alert("你已经完成签到，请不要重复操作！"),this.showControl=!1)}},f={render:function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("div",{directives:[{name:"show",rawName:"v-show",value:e.showControl,expression:"showControl"}]},[t("el-input",{attrs:{placeholder:"password",clearable:"","show-password":""},model:{value:e.password,callback:function(r){e.password=r},expression:"password"}}),t("br"),t("br"),e._v(" "),t("el-button",{staticClass:"btn",attrs:{type:"primary",size:"mini"},on:{click:e.sign}},[e._v("签到")])],1)},staticRenderFns:[]};var R=t("VU/8")(y,f,!1,function(e){t("S0Hs")},null,null).exports,h=(t("tvR6"),t("zL8q")),S=t.n(h);i.default.use(a.a),i.default.use(S.a);var A=new a.a({routes:[{path:"/",name:"HelloWorld",component:R}]});i.default.config.productionTip=!1,new i.default({el:"#app",router:A,components:{App:o},template:"<App/>"})},OJMy:function(e,r,t){var i={};i.web=t("CatX");var n={};n.rapi=t("gYUs"),n.rapi.RapiClient=function(e,r,t){t||(t={}),t.format="binary",this.client_=new i.web.GrpcWebClientBase(t),this.hostname_=e,this.credentials_=r,this.options_=t},n.rapi.RapiPromiseClient=function(e,r,t){t||(t={}),t.format="binary",this.client_=new i.web.GrpcWebClientBase(t),this.hostname_=e,this.credentials_=r,this.options_=t};var o=new i.web.AbstractClientBase.MethodInfo(n.rapi.ApplyRes,function(e){return e.serializeBinary()},n.rapi.ApplyRes.deserializeBinary);n.rapi.RapiClient.prototype.apply=function(e,r,t){return this.client_.rpcCall(this.hostname_+"/rapi.Rapi/apply",e,r||{},o,t)},n.rapi.RapiPromiseClient.prototype.apply=function(e,r){return this.client_.unaryCall(this.hostname_+"/rapi.Rapi/apply",e,r||{},o)};var a=new i.web.AbstractClientBase.MethodInfo(n.rapi.SignRes,function(e){return e.serializeBinary()},n.rapi.SignRes.deserializeBinary);n.rapi.RapiClient.prototype.sign=function(e,r,t){return this.client_.rpcCall(this.hostname_+"/rapi.Rapi/sign",e,r||{},a,t)},n.rapi.RapiPromiseClient.prototype.sign=function(e,r){return this.client_.unaryCall(this.hostname_+"/rapi.Rapi/sign",e,r||{},a)},e.exports=n.rapi},S0Hs:function(e,r){},gYUs:function(e,r,t){var i=t("LkYP"),n=i,o=Function("return this")();n.exportSymbol("proto.rapi.ApplyReq",null,o),n.exportSymbol("proto.rapi.ApplyRes",null,o),n.exportSymbol("proto.rapi.SignReq",null,o),n.exportSymbol("proto.rapi.SignRes",null,o),proto.rapi.ApplyReq=function(e){i.Message.initialize(this,e,0,-1,null,null)},n.inherits(proto.rapi.ApplyReq,i.Message),n.DEBUG&&!COMPILED&&(proto.rapi.ApplyReq.displayName="proto.rapi.ApplyReq"),proto.rapi.ApplyRes=function(e){i.Message.initialize(this,e,0,-1,null,null)},n.inherits(proto.rapi.ApplyRes,i.Message),n.DEBUG&&!COMPILED&&(proto.rapi.ApplyRes.displayName="proto.rapi.ApplyRes"),proto.rapi.SignReq=function(e){i.Message.initialize(this,e,0,-1,null,null)},n.inherits(proto.rapi.SignReq,i.Message),n.DEBUG&&!COMPILED&&(proto.rapi.SignReq.displayName="proto.rapi.SignReq"),proto.rapi.SignRes=function(e){i.Message.initialize(this,e,0,-1,null,null)},n.inherits(proto.rapi.SignRes,i.Message),n.DEBUG&&!COMPILED&&(proto.rapi.SignRes.displayName="proto.rapi.SignRes"),i.Message.GENERATE_TO_OBJECT&&(proto.rapi.ApplyReq.prototype.toObject=function(e){return proto.rapi.ApplyReq.toObject(e,this)},proto.rapi.ApplyReq.toObject=function(e,r){var t={finger:i.Message.getFieldWithDefault(r,1,""),pass:i.Message.getFieldWithDefault(r,2,"")};return e&&(t.$jspbMessageInstance=r),t}),proto.rapi.ApplyReq.deserializeBinary=function(e){var r=new i.BinaryReader(e),t=new proto.rapi.ApplyReq;return proto.rapi.ApplyReq.deserializeBinaryFromReader(t,r)},proto.rapi.ApplyReq.deserializeBinaryFromReader=function(e,r){for(;r.nextField()&&!r.isEndGroup();){switch(r.getFieldNumber()){case 1:var t=r.readString();e.setFinger(t);break;case 2:t=r.readString();e.setPass(t);break;default:r.skipField()}}return e},proto.rapi.ApplyReq.prototype.serializeBinary=function(){var e=new i.BinaryWriter;return proto.rapi.ApplyReq.serializeBinaryToWriter(this,e),e.getResultBuffer()},proto.rapi.ApplyReq.serializeBinaryToWriter=function(e,r){var t=void 0;(t=e.getFinger()).length>0&&r.writeString(1,t),(t=e.getPass()).length>0&&r.writeString(2,t)},proto.rapi.ApplyReq.prototype.getFinger=function(){return i.Message.getFieldWithDefault(this,1,"")},proto.rapi.ApplyReq.prototype.setFinger=function(e){i.Message.setProto3StringField(this,1,e)},proto.rapi.ApplyReq.prototype.getPass=function(){return i.Message.getFieldWithDefault(this,2,"")},proto.rapi.ApplyReq.prototype.setPass=function(e){i.Message.setProto3StringField(this,2,e)},i.Message.GENERATE_TO_OBJECT&&(proto.rapi.ApplyRes.prototype.toObject=function(e){return proto.rapi.ApplyRes.toObject(e,this)},proto.rapi.ApplyRes.toObject=function(e,r){var t={err:i.Message.getFieldWithDefault(r,1,""),addr:i.Message.getFieldWithDefault(r,2,"")};return e&&(t.$jspbMessageInstance=r),t}),proto.rapi.ApplyRes.deserializeBinary=function(e){var r=new i.BinaryReader(e),t=new proto.rapi.ApplyRes;return proto.rapi.ApplyRes.deserializeBinaryFromReader(t,r)},proto.rapi.ApplyRes.deserializeBinaryFromReader=function(e,r){for(;r.nextField()&&!r.isEndGroup();){switch(r.getFieldNumber()){case 1:var t=r.readString();e.setErr(t);break;case 2:t=r.readString();e.setAddr(t);break;default:r.skipField()}}return e},proto.rapi.ApplyRes.prototype.serializeBinary=function(){var e=new i.BinaryWriter;return proto.rapi.ApplyRes.serializeBinaryToWriter(this,e),e.getResultBuffer()},proto.rapi.ApplyRes.serializeBinaryToWriter=function(e,r){var t=void 0;(t=e.getErr()).length>0&&r.writeString(1,t),(t=e.getAddr()).length>0&&r.writeString(2,t)},proto.rapi.ApplyRes.prototype.getErr=function(){return i.Message.getFieldWithDefault(this,1,"")},proto.rapi.ApplyRes.prototype.setErr=function(e){i.Message.setProto3StringField(this,1,e)},proto.rapi.ApplyRes.prototype.getAddr=function(){return i.Message.getFieldWithDefault(this,2,"")},proto.rapi.ApplyRes.prototype.setAddr=function(e){i.Message.setProto3StringField(this,2,e)},i.Message.GENERATE_TO_OBJECT&&(proto.rapi.SignReq.prototype.toObject=function(e){return proto.rapi.SignReq.toObject(e,this)},proto.rapi.SignReq.toObject=function(e,r){var t={finger:i.Message.getFieldWithDefault(r,1,""),addr:i.Message.getFieldWithDefault(r,2,"")};return e&&(t.$jspbMessageInstance=r),t}),proto.rapi.SignReq.deserializeBinary=function(e){var r=new i.BinaryReader(e),t=new proto.rapi.SignReq;return proto.rapi.SignReq.deserializeBinaryFromReader(t,r)},proto.rapi.SignReq.deserializeBinaryFromReader=function(e,r){for(;r.nextField()&&!r.isEndGroup();){switch(r.getFieldNumber()){case 1:var t=r.readString();e.setFinger(t);break;case 2:t=r.readString();e.setAddr(t);break;default:r.skipField()}}return e},proto.rapi.SignReq.prototype.serializeBinary=function(){var e=new i.BinaryWriter;return proto.rapi.SignReq.serializeBinaryToWriter(this,e),e.getResultBuffer()},proto.rapi.SignReq.serializeBinaryToWriter=function(e,r){var t=void 0;(t=e.getFinger()).length>0&&r.writeString(1,t),(t=e.getAddr()).length>0&&r.writeString(2,t)},proto.rapi.SignReq.prototype.getFinger=function(){return i.Message.getFieldWithDefault(this,1,"")},proto.rapi.SignReq.prototype.setFinger=function(e){i.Message.setProto3StringField(this,1,e)},proto.rapi.SignReq.prototype.getAddr=function(){return i.Message.getFieldWithDefault(this,2,"")},proto.rapi.SignReq.prototype.setAddr=function(e){i.Message.setProto3StringField(this,2,e)},i.Message.GENERATE_TO_OBJECT&&(proto.rapi.SignRes.prototype.toObject=function(e){return proto.rapi.SignRes.toObject(e,this)},proto.rapi.SignRes.toObject=function(e,r){var t={err:i.Message.getFieldWithDefault(r,1,"")};return e&&(t.$jspbMessageInstance=r),t}),proto.rapi.SignRes.deserializeBinary=function(e){var r=new i.BinaryReader(e),t=new proto.rapi.SignRes;return proto.rapi.SignRes.deserializeBinaryFromReader(t,r)},proto.rapi.SignRes.deserializeBinaryFromReader=function(e,r){for(;r.nextField()&&!r.isEndGroup();){switch(r.getFieldNumber()){case 1:var t=r.readString();e.setErr(t);break;default:r.skipField()}}return e},proto.rapi.SignRes.prototype.serializeBinary=function(){var e=new i.BinaryWriter;return proto.rapi.SignRes.serializeBinaryToWriter(this,e),e.getResultBuffer()},proto.rapi.SignRes.serializeBinaryToWriter=function(e,r){var t;(t=e.getErr()).length>0&&r.writeString(1,t)},proto.rapi.SignRes.prototype.getErr=function(){return i.Message.getFieldWithDefault(this,1,"")},proto.rapi.SignRes.prototype.setErr=function(e){i.Message.setProto3StringField(this,1,e)},n.object.extend(r,proto.rapi)},gsu9:function(e,r){},tvR6:function(e,r){}},["NHnr"]);
//# sourceMappingURL=app.086d8cdef16ea43a8a3e.js.map