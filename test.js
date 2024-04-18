import { expect } from 'chai';
import { describe, it } from 'mocha';

describe('Compare imports', () => {
  it('different imports', async () => {

    const modA = await import('./a.js');
    const modB = await import('./b.js');

    // throws
    expect({ a: modA }).to.deep.equal({ b: modB});
  });
  it('This would also fail (but never runs)', async () => {

    const foo = Object.create(null, { 
      [Symbol.toStringTag]: { value: 'Foo' }, 
      bing: { get: () => 'bong', enumerable: true  }
    });
    const bar = Object.create(null, { 
      [Symbol.toStringTag]: { value: 'Bar' }, 
      bing: { get: () => 'boing' }
    });

    // throws
    expect({ a: foo }).to.deep.equal({ b: bar });
  });
  it('always true', () => {
    console.log('This will never run');
  });
});