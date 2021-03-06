# Packs

## Problem

Imagine for a moment that one of our product lines ships in various pack sizes:

- 250 Items
- 500 Items
- 1000 Items
- 2000 Items
- 5000 Items

Our customers can order any number of these items through our website, but they will always only be given complete packs.

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfil each order.

For instance;

1 item - 1 x 250 \
250 items - 1 x 250 \
251 items - 1 x 500 \
501 items - 1 x 500, 2 x 250 \
12001 items - 2 x 5000, 1 x 2000, 1 x 250

## API

The service is deployed on AWS Lambda using Serverless. It can be deployed by running `./deploy.sh`. The function can be invoked using the following cURL;

```
curl -X POST \
-d '{"count": 666, "pack_sizes": [250,500, 1000, 2000, 5000]}' \
https://gi891qg1u8.execute-api.eu-west-1.amazonaws.com/dev/calculate
```

Should receive the following JSON response;

```
{"packs":[{"size":250,"count":1},{"size":500,"count":1}]}
```
