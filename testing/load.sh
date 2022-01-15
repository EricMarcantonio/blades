k6 run --vus 100 --duration 1s load.js --summary-export load_100_1s.json
k6 run --vus 250 --duration 1s load.js --summary-export load_250_1s.json
k6 run --vus 500 --duration 1s load.js --summary-export load_500_1s.json