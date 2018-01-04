import { assert } from 'chai';
import rewiremock from 'rewiremock';

import * as exports from 'superagent';
import * as UnitUnderTest from '../src/js/api';

rewiremock.around(() => import('superagent')).then(mockedModule => ({
  post: {
    send: () => ({
      end: (cb) => {
        cb(null, JSON.stringify({Url: 'localhost:8080/eZ8dK'}));
      }
    })
  }
}));

describe('api', () => {
  it('calls post', () => {
    UnitUnderTest.generateCode('https://google.com');
  });
});