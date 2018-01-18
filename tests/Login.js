
module.exports = {
    'Test Login With Valid Details Succeeds' : function (browser) {
      browser
        .url('http://localhost:8000')
        .waitForElementVisible('#login-link', 1000)
        .click('#login-link')
        .waitForElementVisible('input[type=email]', 1000)
        .setValue('input[type=email]', 'testaccount@test.com')
        .setValue('input[type=password]', 'SuperSecretPassword')
        .waitForElementVisible('button[id=login-button]', 1000)
        .click('button[id=login-button]')
        .pause(1000)
        .assert.containsText('body', 'Login Success')
        .end();
    },
    'Test Login With Incorrect Details Fails' : function (browser) {
      browser
      .url('http://localhost:8000')
      .waitForElementVisible('#login-link', 1000)
      .click('#login-link')
      .waitForElementVisible('input[type=email]', 1000)
      .setValue('input[type=email]', 'blah@test.com')
      .setValue('input[type=password]', 'password')
      .waitForElementVisible('button[id=login-button]', 1000)
      .click('button[id=login-button]')
      .pause(1000)
      .assert.containsText('body', 'Login Failed')
      .end();
    }
  };