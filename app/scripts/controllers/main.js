'use strict';

/**
 * @ngdoc function
 * @name myjsfiddleApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the myjsfiddleApp
 */
angular.module('myjsfiddleApp')
  .controller('MainCtrl', function ($scope,$state) {
    $state.go('main.fiddle');
  });
