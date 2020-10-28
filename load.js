//k6 run -d 10s --vus 100 single-request.js

import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';
export const requests = new Counter('http_reqs');

export default function () {

    const res = http.get('http://127.0.0.1:5000/v1/1');

    const checkRes = check(res, {
        'status is 200': (r) => r.status === 200,
    });
}