it("loads", () => {
  cy.visit("/");
  cy.wait(1000);
});

it("filters logs based on node_id(invalid)", () => {
  cy.visit("/report");
  cy.get("[data-cy=node_id]").type("{enter}");
  cy.get("[class=notification-content]").should(
    "contain.text",
    "couldnt create proctor"
  );
  cy.wait(2000);
  // cy.screenshot()
});
