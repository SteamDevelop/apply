let utils = {
  apply: function (pwd) {
    const {ReqData, ResData} = require('./hi_pb.js');
    const {HiDotClient} = require('./hi_grpc_web_pb.js');

    let rpcweb = new HiDotClient('http://localhost:6868');

    let request = new ReqData();
    request.setName('http hi client');

    rpcweb.hi(request, pwd, (err, response) => {

      console.log("Instead of memory. ");
      if (response) {
        console.log(response.getName());
      }else{
        console.log(err);
      }
    });

    let rpcwebs = new HiDotClient('https://localhost:6868');
    rpcwebs.hi(request, {}, (err, response) => {
      console.log("http");
      if (response) {
        console.log(response.getName());
      }else{
        console.log(err);
      }
    });
  }
};

export {utils};
