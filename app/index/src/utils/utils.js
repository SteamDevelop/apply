

import { RapiPromiseClient } from './apply_grpc_web_pb'
import { ApplyReq, SignReq } from './apply_pb'

export class ApplyService {
  static apply (params) {
    let request = new ApplyReq()
    let loginService = new RapiPromiseClient("https://127.0.0.1:443/rapi", null, null)
    request.setPass(params.password)
    request.setFinger(params.finger)
    return loginService.apply(request, null)
  }

  static sign(params) {
    let request = new SignReq()
    let loginService = new RapiPromiseClient("https://localhost:443/rapi", null, null)
    // request.setPass(params.password)
    request.setFinger(params.finger)
    return loginService.sign(request, null)
  }
}
