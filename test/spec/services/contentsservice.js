'use strict';

describe('Service: contentsService', function () {

  // load the service's module
  beforeEach(module('myjsfiddleApp'));

  // instantiate service
  var contentsService;
  beforeEach(inject(function (_contentsService_) {
    contentsService = _contentsService_;
  }));

  it('should do something', function () {
    expect(!!contentsService).toBe(true);
  });

});
