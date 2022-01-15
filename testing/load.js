import http from "k6/http";


export default function () {
    let reqs = [{
        method: 'POST',
        url: 'http://blades.ericmarcantonio.com:3000/gql',
        body: '{"query":"query {\\n    products{\\n        id\\n    }\\n}"}',
    },
    {
        method: 'POST',
        url: 'http://blades.ericmarcantonio.com:3000/gql',
        body: '{"query":"query {\n    product(id:1){\n        name\n        price\n    }\n}","variables":{"id":1}}',
    }, {
        method: 'POST',
        url: 'http://blades.ericmarcantonio.com:3000/gql',
        body: '{"query":"mutation {\n    createProduct(name: \"A New Skate\", price: 49.99, units: 3){\n        id\n        units\n        name\n        price\n        modified_date\n        added_date\n        is_active\n        \n    }\n}","variables":{}}',

    },
    {
        method: 'POST',
        url: 'http://blades.ericmarcantonio.com:3000/gql',
        body: '{"query":"mutation {\n    createProduct(name: \"A New Skate\", price: 49.99){\n        id\n        units\n        name\n        price\n        modified_date\n        added_date\n        is_active\n        \n    }\n}","variables":{}}',

    },
    {
        method: 'POST',
        url: 'http://blades.ericmarcantonio.com:3000/gql',
        body: '{"query":"mutation {\n    deactivateProduct(id: 5){\n        id\n        is_active\n    }\n}","variables":{}}',

    }]

    http.batch(reqs)
}






