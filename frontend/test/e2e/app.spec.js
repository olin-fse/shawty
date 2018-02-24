const chai = require('chai');
const expect = chai.expect;

const config = require('../../config')(process.env.NODE_ENV);

describe('shawty', function() {
  beforeEach(function () {
    browser.url('http://localhost:9001');
  });

  it('renders app', async function() {
    const res = await browser.element('.App');
    expect(res).to.not.be.null;
  });

  it('app flow works', async function() {
    browser.setValue('input[type="url"]', 'https://google.com');
    browser.click('input[type="submit"]');
    browser.waitForText('.App-result', 1000);
    const url = browser.getText('.App-result');
    expect(url).to.have.length(config.staticEndpoint.length + 6);

    browser.url(url);
    const title = browser.getTitle();
    expect(title).to.equal('Google');
  });
});