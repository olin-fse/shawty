import request from 'superagent';

import * as UnitUnderTest from '../src/js/api';

test('api', () => {
  const urlCode = "abcde";
  const spy = jest.spyOn(request, 'post').mockImplementation(() => (
    {
      send: () => ({
        end: (cb) => {
          return cb(null, { text: JSON.stringify({ Url: urlCode }) });
        }
      })
    }
  ));

  UnitUnderTest.generateCode('https://google.com', (code) => {
    expect(spy).toHaveBeenCalled();
    expect(code).toBe(urlCode);
  });

  spy.mockReset();
  spy.mockRestore();
});