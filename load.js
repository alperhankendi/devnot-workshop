//k6 run -d 10s --vus 100 single-request.js

import http from 'k6/http';
import { sleep, check } from 'k6';
import { Counter } from 'k6/metrics';
export default function () {

    const res = http.get(`http://127.0.0.1:5000/api/v1/${__ENV.ID}`);

    const checkRes = check(res, {
        'status is 200': (r) => r.status === 200,
    });

}