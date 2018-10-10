const app = angular.module("short-url", []);

app.controller("mainCtrl", function ($scope, $window, $http) {
    $scope.url = null;
    $scope.longUrl = null;
    $scope.shortUrl = null;
    $scope.alerbox = 'd-none';
    $scope.host = $window.location.host;
    $scope.shortenUrl = () => {
        if (!$scope.url) {
            return;
        }

        $http
            .post("/shorten", {
                url: $scope.url
            })
            .then(response => {
                $scope.longUrl = response.data.long_url;
                $scope.shortUrl = response.data.link;
                $scope.alertbox = 'd-block';
            }, error => {
                console.log(error)
            });

        // reset the field
        $scope.url = null;
    };

    $scope.closeAlertbox = () => {
        $scope.alertbox = 'd-none';
    };

    $scope.copy = () => {
        let text = $window.location.host + $scope.shortUrl;
        copyToClipboard(text)
    };
});

function copyToClipboard(text) {
    const aux = document.createElement("input");    // create a "hidden" input
    aux.setAttribute("value", text); // assign it the value of the specified element
    document.body.appendChild(aux);                 // append it to the body
    aux.select();                                   // highlight its content
    document.execCommand("copy");                   // copy the highlighted text
    document.body.removeChild(aux);                 // remove it from the body
}
