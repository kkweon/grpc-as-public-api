import { HelloClient } from './proto/Hello_serviceServiceClientPb'
import { HelloRequest, HelloResponse } from './proto/hello_service_pb'

const client = new HelloClient('https://k8s.kkweon.dev/envoy')
const request = new HelloRequest()
request.setName('Mo Kweon')

client.say(request, null, (err, resp: HelloResponse) => {
	if (err) {
		console.error(err)
	}
	console.log(`Received a message from the server => ${resp.getMessage()}`)
	console.log({ resp })
})
