const fastify = require("fastify")({ logger: false });

fastify.get("/", async function (request, reply) {
  return { hello: "world" }
});

const start = async () => {
  try {
    await fastify.listen({ port: 3001 })
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}
start()
