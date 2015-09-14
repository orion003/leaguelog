angular.module('rlApp').config(['$stateProvider', '$urlRouterProvider', '$locationProvider',
    function($stateProvider, $urlRouterProvider, $locationProvider) {
        $stateProvider
            .state('landing', {
                url: '',
                abstract: true,
                templateUrl: '/app/components/landing/landing.html'
            })
            .state('landing.form', {
                url: '/',
                resolve: {
                    user: function() { return {}; }
                },
                templateUrl: '/app/components/landing/landing.form.html',
                controller: 'LandingController',
                controllerAs: 'landing'
            })
            .state('landing.success', {
                templateUrl: '/app/components/landing/landing.success.html'
            })
            .state('landing.duplicate', {
                templateUrl: '/app/components/landing/landing.duplicate.html'
            });
            
        $urlRouterProvider.otherwise('/');
        
        $locationProvider.html5Mode(true);
    }]);