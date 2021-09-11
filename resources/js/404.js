$(document).ready(function () {
    var inputReady = true;
    var input = $('.not-found-input');
    input.focus();
    $('.container').on('click', function (e) {
        input.focus();
    });

    input.on('keyup', function (e) {
        $('.new-output').text(input.val());
    });

    $('.four-oh-four-form').on('submit', function (e) {
        e.preventDefault();
        var command = $(this).children($('.not-found-input')).val().toLowerCase()
        switch (command) {
            case 'back':
                var input = $('.not-found-input');
                $('.new-output').removeClass('new-output');
                input.val('');
                $('.terminal').append(
                    '<p class="prompt">OK. Return to the previous page in 5 seconds </p>' +
                    '<p class="prompt">Bye ~~ </p>' +
                    '<p class="prompt output new-output"></p>'
                );

                setTimeout(function () {
                    history.back();
                }, 3000);
                break
            case 'help':
                showHelp();
                break
            default:
                resetForm();
                break;
        }
    });

    function resetForm(withKittens) {
        var message = "Sorry that command is not recognized."
        var input = $('.not-found-input');

        if (withKittens) {
            $('.mingo').removeClass('mingo');
            message = "Huzzzzzah Mingo!"
        }

        $('.new-output').removeClass('new-output');
        input.val('');
        $('.terminal').append('<p class="prompt">' + message + '</p><p class="prompt output new-output"></p>');
    }

    function showHelp() {
        $('.terminal').append(
            "<div class='mingo'>" +
            "<p class='prompt'>	                             ,----,         ,----,                                          ,---,</p>" +
            "<p class='prompt'>       ,--.                ,/   .`|       ,/   .`|                     ,--.              ,`--.' |</p>" +
            "<p class='prompt'>   ,--/  /|    ,---,     ,`   .'  :     ,`   .'  :     ,---,.        ,--.'|   .--.--.    |   :  :</p>" +
            "<p class='prompt'>,---,': / ' ,`--.' |   ;    ;     /   ;    ;     /   ,'  .' |    ,--,:  : |  /  /    '.  '   '  ;</p>" +
            "<p class='prompt'>:   : '/ /  |   :  : .'___,/    ,'  .'___,/    ,'  ,---.'   | ,`--.'`|  ' : |  :  /`. /  |   |  |</p>" +
            "<p class='prompt'>|   '   ,   :   |  ' |    :     |   |    :     |   |   |   .' |   :  :  | | ;  |  |--`   '   :  ;</p>" +
            "<p class='prompt'>'   |  /    |   :  | ;    |.';  ;   ;    |.';  ;   :   :  |-, :   |   \\ | : |  :  ;_     |   |  '</p>" +
            "<p class='prompt'>|   ;  ;    '   '  ; `----'  |  |   `----'  |  |   :   |  ;/| |   : '  '; |  \\  \\    `.  '   :  |</p>" +
            "<p class='prompt'>:   '   \\   |   |  |     '   :  ;       '   :  ;   |   :   .' '   ' ;.    ;   `----.   \\ ;   |  ;</p>" +
            "<p class='prompt'>'   : |.  \\ |   |  '     '   :  |       '   :  |   '   :  ;/| '   : |  ; .'  /  /`--'  /  `--..`;  </p>" +
            "<p class='prompt'>|   | '_\\.' '   :  |     ;   |.'        ;   |.'    |   |    \\ |   | '`--'   '--'.     /  .--,_   </p>" +
            "<p class='prompt'>'   : |     ;   |.'      '---'          '---'      |   :   .' '   : |         `--'---'   |    |`.  </p>" +
            "<p class='prompt'>;   |,'     '---'                                  |   | ,'   ;   |.'                    `-- -`, ; </p>" +
            "<p class='prompt'>'---'                                              `----'     '---'                        '---`'</p>" +
            "<p class='prompt'>                                                              </p>" +
            "</div>");


        var lines = $('.mingo p');
        $.each(lines, function (index, line) {
            setTimeout(function () {
                $(line).css({
                    "opacity": 1
                });

                textEffect($(line))
            }, index * 100);
        });
    }

    function textEffect(line) {
        var alpha = [';', '.', ',', ':', ';', '~', '`'];
        var animationSpeed = 10;
        var index = 0;
        var string = line.text();
        var splitString = string.split("");
        var copyString = splitString.slice(0);

        var emptyString = copyString.map(function (el) {
            return [alpha[Math.floor(Math.random() * (alpha.length))], index++];
        })

        emptyString = shuffle(emptyString);

        $.each(copyString, function (i, el) {
            var newChar = emptyString[i];
            toUnderscore(copyString, line, newChar);

            setTimeout(function () {
                fromUnderscore(copyString, splitString, newChar, line);
            }, i * animationSpeed);
        })
    }

    function toUnderscore(copyString, line, newChar) {
        copyString[newChar[1]] = newChar[0];
        line.text(copyString.join(''));
    }

    function fromUnderscore(copyString, splitString, newChar, line) {
        copyString[newChar[1]] = splitString[newChar[1]];
        line.text(copyString.join(""));
    }


    function shuffle(o) {
        for (var j, x, i = o.length; i; j = Math.floor(Math.random() * i), x = o[--i], o[i] = o[j], o[j] = x) ;
        return o;
    };
});