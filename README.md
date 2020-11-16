README

#### Setup

```
k6 run --duration 60s --vus 150 load-get.js -e ID=7lf1ly4khsnhzfqsntec

running (1m00.1s), 000/150 VUs, 3126018 complete and 0 interrupted iterations
default ✓ [======================================] 150 VUs  1m0s
    ✗ status is 200
     ↳  99% — ✓ 3125731 / ✗ 287
```

