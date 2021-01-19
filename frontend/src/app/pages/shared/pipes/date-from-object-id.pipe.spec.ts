import { DateFromObjectIdPipe } from './date-from-object-id.pipe';

describe('DateFromObjectIdPipe', () => {
  it('create an instance', () => {
    const pipe = new DateFromObjectIdPipe();
    expect(pipe).toBeTruthy();
  });
});
