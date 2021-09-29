/// <reference types="cypress" />

let apiEndPoint = "localhost:9090";

describe("tests for comand dispatcher API", () => {
  it("verify request returns JSON", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/configuration/`,

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
      url: `${apiEndPoint}/configuration/`,

      body: { limit: 1 },
    })
      .its("status")
      .should("be.equal", 200);
  });

  it("verify the request returns 3 items", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/configuration/`,

      body: { limit: 3 },
    })
      .its("body.data")
      .should("have.length", 3);
  });

  it("verify the request returns 3 items (invalid) ", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/configuration`,

      body: { limit: "dd" },
      failOnStatusCode: false,
    }).should((response) => {
      cy.log(JSON.stringify(response.body));
    });
  });

  it("verify an array of config reports can be fetched with a valid limit", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/configuration/config-reports`,

      body: { limit: "4" },
    })
      .then((response) => {
        expect(response.status).to.eq(200);
      })
      .its("body.data")
      .should("have.length", 4);
  });

  it("verify a valid configuration profile could be added", () => {
    cy.request({
      method: "POST",
      url: `${apiEndPoint}/configuration/add-config-profile`,

      body: {
        jump_spacing: 2,
        replication_count: 3,
      },
    }).then((response) => {
      expect(response.status).to.eq(200);
      cy.log(JSON.stringify(response.body));
    });
  });

  it("verify the currently active configuration profile can be retrived", () => {
    cy.request({
      method: "GET",
      url: `${apiEndPoint}/configuration/current`,
      failOnStatusCode: false,
    }).then((response) => {
      expect(response.status).to.eq(200);
      cy.log(JSON.stringify(response.body));
    });
  });

  it("verify the jump space balancing for the nodes can be changed successfully", () => {
    cy.request({
      method: "POST",
      url: `${apiEndPoint}/configuration/set-jump-space`,

      body: {
        new_jump_space: 3,
      },
    }).then((response) => {
      expect(response.status).to.eq(200);
      cy.log(JSON.stringify(response.body.message));
    });
  });

  it("verify the replication level for the nodes can be changed successfully", () => {
    cy.request({
      method: "POST",
      url: `${apiEndPoint}/configuration/set-replication-level`,

      body: {
        new_replication_level: 3,
      },
    }).then((response) => {
      expect(response.status).to.eq(200);
      cy.log(JSON.stringify(response.body.message));
    });
  });

  it("verify a configuration profile could be deleted with a valid id", () => {
    cy.request({
      method: "DELETE",
      url: `${apiEndPoint}/configuration/delete-config-profile`,

      body: {
        configuration_id: "61544c2c766ec3cee4b7f0b5",
      },
      failOnStatusCode: false,
    }).then((response) => {
      cy.log(JSON.stringify(response.body));
      expect(response.status).to.eq(200);
    });
  });
});
