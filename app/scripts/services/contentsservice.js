'use strict';

/**
 * @ngdoc service
 * @name myjsfiddleApp.contentsService
 * @description
 * # contentsService
 * Service in the myjsfiddleApp.
 */
angular.module('myjsfiddleApp')
  .service('ContentsService', function contentsService($http) {
    var ContentsService = function(){
    };
    
    ContentsService.prototype.find = function(themeId,id){
      return $http.get('/app/theme/' + themeId +'/' + id);
    };

    ContentsService.prototype.createTheme = function(theme){
      return $http.post('/app/theme',theme);
    };

    return new ContentsService();

  });
