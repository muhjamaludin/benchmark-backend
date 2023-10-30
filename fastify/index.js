const fastify = require("fastify")({ logger: false });

fastify.get("/", function handler(request, reply) {
  reply.send({ hello: "world" });
});

fastify.listen({ port: 3001 }, (err) => {
  if (err) {
    fastify.log.error(err);
    process.exit(1);
  }
});
