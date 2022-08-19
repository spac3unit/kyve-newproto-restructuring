
This query servers all queries for the following modules:
bundles, pool, delegation, stakers
as a lot of request require composition of multiple module-data
everything will be served from one single query module.

Therefore, the queries are allowed to have the package prefixes
of these modules in their query url
The named package do not implement own requires (expect for params)


For simplicity all queries and their objects are in the corresponding
proto files