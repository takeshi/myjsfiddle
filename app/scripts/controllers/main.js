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

    $scope.Mode = {
      perspective:'css'
    };

    $scope.cssMode = function(){
      $scope.Mode.perspective = 'css';
    };

    $scope.directiveMode = function(){
      $scope.Mode.perspective = 'directive';
    };

    $scope.editorHtmlOptions = {
      lineWrapping : true,
      lineNumbers: true,
      keyMap: "sublime",
      autoCloseBrackets: true,
      matchBrackets: true,
      showCursorWhenSelecting: true,
      theme: "monokai",
      mode: 'xml'
    };

    $scope.editorSassOptions = {
      lineWrapping : true,
      lineNumbers: true,
      keyMap: "sublime",
      autoCloseBrackets: true,
      matchBrackets: true,
      showCursorWhenSelecting: true,
      theme: "monokai",
      mode: 'sass'
    };

    $scope.editorJsOptions = {
      lineWrapping : true,
      lineNumbers: true,
      keyMap: "sublime",
      autoCloseBrackets: true,
      matchBrackets: true,
      showCursorWhenSelecting: true,
      theme: "monokai",
      mode: 'javascript'
    };

  });
