const fastify = require("fastify")({ logger: false });

fastify.get("/", function (request, reply) {
  return { hello: "world" }
});

fastify.listen({ port: 3001 }, (err) => {
  if (err) {
    fastify.log.error(err);
    process.exit(1);
  }
});
