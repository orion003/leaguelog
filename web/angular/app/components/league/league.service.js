angular.module('league')
    .service('LeagueService', ['$http', function($http) {
            
            this.getLeague = function(leagueId) {
                return $http.get('/api/league/' + leagueId);
            };
            
            this.getStandings = function(leagueId) {
                return $http.get('/api/league/' + leagueId + '/standings');
            };
            
            this.getSchedules = function(leagueId) {
                return $http.get('/api/league/' + leagueId + '/schedule');
            };
            
    }]);