// Auto-generated by the postman-to-k6 converter

import "./libs/shim/core.js";
import "./libs/shim/expect.js";
import "./libs/shim/urijs.js";

export let options = { maxRedirects: 4 };

const Request = Symbol.for("request");
postman[Symbol.for("initial")]({
  options,
  collection: {
    url: "http://blades.ericmarcantonio.com:3000/gql"
  }
});

export default function() {
  postman[Request]({
    name: "GetAllProducts",
    id: "bb8f4c78-a0de-4277-9b9c-76ae4979b450",
    method: "POST",
    address: "{{url}}",
    data: '{"query":"query {\\n    products{\\n        id\\n    }\\n}"}',
    post(response) {
      pm.test("Contains products", function() {
        var jsonData = pm.response.json();
        pm.expect(jsonData.data.products).is.not.null;
      });
    }
  });

  postman[Request]({
    name: "GetProductById",
    id: "946560af-92b0-449b-aa1d-742689e18f73",
    method: "POST",
    address: "{{url}}",
    data:
      '{"query":"query {\\n    product(id:1){\\n        name\\n        price\\n    }\\n}","variables":"{\\n    \\"id\\": 1\\n}"}',
    post(response) {
      pm.test("Contains a product fields", function() {
        var jsonData = pm.response.json();
        pm.expect(jsonData.data.product.name).to.eql(
          "Bauer Supreme Ultrasonic Skates"
        );
        pm.expect(jsonData.data.product.price).to.eql(599.99);
      });
    }
  });

  postman[Request]({
    name: "CreateAProduct",
    id: "e46e45f6-68e1-4459-9975-ea557f3ddd97",
    method: "POST",
    address: "{{url}}",
    data:
      '{"query":"mutation {\\n    createProduct(name: \\"A New Skate\\", price: 49.99, units: 0){\\n        id\\n        units\\n        name\\n        price\\n        modified_date\\n        added_date\\n        is_active\\n        \\n    }\\n}"}',
    post(response) {
      pm.test("Create a Product", function() {
        var jsonData = pm.response.json();
        pm.expect(jsonData.data.createProduct.added_date).not.null;
        pm.expect(jsonData.data.createProduct.id).not.null;
        pm.expect(jsonData.data.createProduct.is_active).eq("yes");
        pm.expect(jsonData.data.createProduct.modifed_date).not.null;
        pm.expect(jsonData.data.createProduct.name).eq("A New Skate");
        pm.expect(jsonData.data.createProduct.units).eq(0);
      });
    }
  });

  postman[Request]({
    name: "UpdateAProduct",
    id: "c61aac1b-a773-481b-9e28-de626279fa52",
    method: "POST",
    address: "{{url}}",
    data:
      '{"query":"mutation {\\n    updateProduct(id: 1, name: \\"A New Skate Updated\\", price: 40.99, units: 12){\\n        id\\n        units\\n        name\\n        price\\n        modified_date\\n        added_date\\n        is_active\\n        \\n    }\\n}"}',
    post(response) {
      pm.test("Update a Product", function() {
        var jsonData = pm.response.json();
        pm.expect(jsonData.data.updateProduct.added_date).not.null;
        pm.expect(jsonData.data.updateProduct.id).is(1);
        pm.expect(jsonData.data.updateProduct.is_active).eq("yes");
        pm.expect(jsonData.data.updateProduct.modifed_date).not.null;
        pm.expect(jsonData.data.updateProduct.name).eq("A New Skate Updated");
        pm.expect(jsonData.data.updateProduct.units).eq(12);
      });
    }
  });

  postman[Request]({
    name: "DeleteAProduct",
    id: "3e77e775-04dd-4925-a0dd-3e0f325d7d7e",
    method: "POST",
    address: "{{url}}",
    data:
      '{"query":"mutation {\\n    deactivateProduct(id: 5){\\n        id\\n        is_active\\n    }\\n}"}',
    post(response) {
      pm.test("Update a Product", function() {
        var jsonData = pm.response.json();
        pm.expect(jsonData.data.deactivateProduct.id).is(5);
        pm.expect(jsonData.data.deactivateProduct.is_active).eq("no");
      });
    }
  });
}
