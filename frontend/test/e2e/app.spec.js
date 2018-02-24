const chai = require('chai');
const expect = chai.expect;

const getConfig = require('../../config');
const config = getConfig(process.env.NODE_ENV);

describe('shawty', function() {
  beforeEach(function () {
    browser.url('http://localhost:9001');
  });

  it('renders app', async function() {
    const res = await browser.element('.App');
    expect(res).to.not.be.null;
  });

  it('submits link and receives code', async function() {
    browser.setValue('input[type="url"]', 'https://google.com');
    browser.click('input[type="submit"]');
    browser.waitForText('.App-result', 1000);
    const text = browser.getText('.App-result');
    expect(text).to.have.length(config.staticEndpoint.length + 6);
  });
});