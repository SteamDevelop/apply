
import { RapiPromiseClient } from './apply_grpc_web_pb'
import { ApplyReq, SignReq } from './apply_pb'

export class ApplyService {
    static apply (params) {
        let request = new ApplyReq()
        let loginService = new RapiPromiseClient("https://..../rapi", null, null)
        request.setPassword(params.password)
        request.setFinger(params.finger)
        return loginService.login(request, null)
    }

    static sign(params) {
        let request = new SignReq()
        let loginService = new RapiPromiseClient("https://..../rapi", null, null)
        request.setPassword(params.password)
        request.setFinger(params.finger)
        return loginService.sign(request, null)
    }
}

ApplyService.apply({password:"sdf", finger:"sdfd"}).then(data => {

}).catch((err) => {

})
