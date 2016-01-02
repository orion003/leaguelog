angular.module("auth", []);
angular.module("home", ['auth']);
angular.module("league", []);

angular.module('rlApp', [
    'ui.router',
    'angularytics',
    'ngStorage',
    'auth',
    'home',
    'league']);
