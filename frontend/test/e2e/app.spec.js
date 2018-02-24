const chai = require('chai');
const chaiWebdriver = require('chai-webdriverio').default;
chai.use(chaiWebdriver(browser));
const expect = chai.expect;

describe('shawty', function() {
  beforeEach(function () {
    browser.url('http://localhost:9001');
  });

  it('renders app', async function() {
    expect('.App').to.be.visible();
  });

  it('app flow works', async function() {
    browser.setValue('input[type="url"]', 'https://google.com');
    browser.click('input[type="submit"]');
    browser.waitForText('.App-result', 1000);

    const regex = /http:\/\/localhost:9001\/[a-zA-Z0-9]{5}/;
    expect('.App-result').to.have.text(regex);

    const url = browser.getText('.App-result');
    browser.url(url);
    expect(browser.getTitle()).to.equal('Google');
  });
});