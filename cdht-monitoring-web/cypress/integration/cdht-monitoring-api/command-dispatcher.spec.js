/// <reference types="cypress" />

let apiEndPoint = "localhost:9090";

describe("tests for comand dispatcher API", () => {
  it("verify request returns JSON", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/all`,

      body: { limit: 1 },
    })
      .should((response) => {
        expect(response.status).to.eq(200);
      })
      .its("headers")
      .its("content-type")
      .should("include", "application/json");
  });

  it("verify the request returns the correct status code for a request", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/all`,

      body: { limit: 5 },
    })
      .its("status")
      .should("be.equal", 200);
  });

  it("verify the request returns 10 items", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/all`,

      body: { limit: 10 },
    })
      .its("body.data")
      .should("have.length", 10);
  });

  it("verify the request returns 10 items (invalid) ", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/all`,

      body: { limit: "dd" },
      failOnStatusCode: false,
    }).should((response) => {
      cy.log(JSON.stringify(response.body));
    });
  });

  it("verify the a single command can be retrived with a valid id", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/single-command`,

      body: { command_id: "6141b792c0b86338549b2c6a" },
    }).then((response) => {
      expect(response.status).to.eq(200);
    });
  });

  it("verify the a single command can not be retrived with a invalid id", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/single-command`,

      body: { command_id: 4545 },
      failOnStatusCode: false,
    }).then((response) => {
      expect(response.status).to.eq(500);
      cy.log(JSON.stringify(response.body));
    });
  });

  it("verify an array of results for commands can be fetched with a valid limit", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/command-dispatcher/result`,

      body: { limit: 4 },
    })
      .then((response) => {
        expect(response.status).to.eq(200);
      })
      .its("body.data")
      .should("have.length", 4);
  });

  it("verify a valid command response can be posed on the databse with a valid object", () => {
    cy.request({
      method: "POST",
      url: `${apiEndPoint}/command-dispatcher/result`,

      body: {
        type: "COMMAND_TYPE_PING",
        operation_id: "6152e11f6b83df987dbee4ff",
        operation_status: "",
        node_id: "",
        node_address: "",
        body: { node_id: null },
      },
    }).then((response) => {
      expect(response.status).to.eq(200);
      cy.log(JSON.stringify(response.body));
    });
  });

  it("verify a valid command response can not  be posed on the databse with a invalid object", () => {
    cy.request({
      method: "POST",
      url: `${apiEndPoint}/command-dispatcher/result`,

      body: {},
      failOnStatusCode: false,
    }).then((response) => {
      expect(response.status).to.eq(500);
      cy.log(JSON.stringify(response.body));
    });
  });
});
