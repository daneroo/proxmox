'use strict'
const pulumi = require('@pulumi/pulumi')
const docker = require('@pulumi/docker')

const network = new docker.Network('hello_default')

const natsImage = new docker.RemoteImage('nats-image', {
  name: 'synadia/nats-server:nightly',
  keepLocally: true // don't delete the image from the local machine when deleting this resource.
})

const natsContainer = new docker.Container('nats', {
  image: natsImage.name,
  networksAdvanced: [{ name: network.name }],
  restart: 'unless-stopped',
  ports: [{
    internal: 8222,
    external: 28222
  }]
})

const natsEndpoint = natsContainer.ports.apply(p => `http://localhost:${p[0].external}`)
exports.appName = natsContainer.name
exports.host = natsEndpoint
