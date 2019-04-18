<!DOCTYPE html>
<html lang="en">

<head>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="A personal webpage for Keith Hand, asoftware engineer and web developer based out of Oklahoma City, Oklahoma.">
    <meta name="author" content="keith.hand">
    <title>Keith Hand - Developer</title>

    <!-- Bootstrap CSS -->
    <link href="css/bootstrap.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link href="css/main.css" rel="stylesheet">

    <!-- Custom Fonts -->
    <link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.8.1/css/all.css" integrity="sha384-50oBUHEmvpQ+1lW4y57PTFmhCaXp0ML5d60M1M7uH2+nqUivzIebhndOJK28anvf" crossorigin="anonymous">
    <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,700,300italic,400italic,700italic" rel="stylesheet" type="text/css">

</head>

<body>

    <!-- Navigation -->
    <a id="menu-toggle" href="#" class="btn btn-dark btn-lg toggle"><i class="fa fa-bars"></i></a>
    <nav id="sidebar-wrapper">
        <ul class="sidebar-nav">
            <a id="menu-close" href="#" class="btn btn-light btn-lg pull-right toggle"><i class="fa fa-times"></i></a>
            <li class="sidebar-brand">
                <a href="#top" onclick=$("#menu-close").click();>Start Bootstrap</a>
            </li>
            <li>
                <a href="#top" onclick=$("#menu-close").click();>Home</a>
            </li>
            <li>
                <a href="#about" onclick=$("#menu-close").click();>About</a>
            </li>
            <li>
                <a href="#services" onclick=$("#menu-close").click();>Skills</a>
            </li>
            <li>
                <a href="#portfolio" onclick=$("#menu-close").click();>Portfolio</a>
            </li>
            <li>
                <a href="#contact" onclick=$("#menu-close").click();>Contact</a>
            </li>
        </ul>
    </nav>

    <!-- Header -->
    <header id="top" class="header bg-primary">
        <div class="text-vertical-center">
            <h1>Keith Hand</h1>
            <h3>Software Developer</h3>
            <br>
            <a href="#about" class="btn btn-light btn-lg">Who am I?</a>
        </div>
    </header>

    <!-- About -->
    <section id="about" class="about">
        <div class="container">
            <div class="row">
                <div class="col-lg-12 text-center">
                    <h2>Hi! My name is Keith and I like to write code.</h2>
                    <p class="lead">I enjoy taking on any project, big or small. Whether you need a personal webpage you have the next big AAA iOS title, or you need someone to manage your webserver, I want to be the guy to help make your dreams become a reality.</p>
                </div>
            </div>
            <!-- /.row -->
        </div>
        <!-- /.container -->
    </section>

    <!-- Services -->
    <!-- The circle icons use Font Awesome's stacked icon classes. For more information, visit http://fontawesome.io/examples/ -->
    <section id="services" class="services bg-primary">
        <div class="container">
            <div class="row text-center">
                <div class="col-lg-10 col-lg-offset-1">
                    <h2>My Knowledge Base</h2>
                    <hr class="small">
                    <!-- row 1 (Webpages) -->
                    <div class="row">
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-html5 fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>HTML</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-php fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>PHP</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-js-square fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>JavaScript</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-css3-alt fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>CSS</strong>
                                </h4>
                            </div>
                        </div>
                    </div>
                    <!-- /.row (nested) -->
                    <!-- row 2 (Languages) -->
                    <div class="row">
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fas fa-laptop-code fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>C++</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-python fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Python</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="far fa-gem fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Ruby</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-java fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Java</strong>
                                </h4>
                            </div>
                        </div>
                    </div>
                    <!-- /.row (nested) -->
                    <!-- row 3 (Programs) -->
                    <div class="row">
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-git fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Git</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-docker fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Docker</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fas fa-database fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>SQL</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="far fa-file-alt fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>LaTeX</strong>
                                </h4>
                            </div>
                        </div>
                    </div>
                    <!-- /.row (nested) -->
                    <!-- row 4 (Operating System) -->
                    <div class="row">
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fa fa-terminal fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>sh</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-linux fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Linux</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-windows fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>Windows</strong>
                                </h4>
                            </div>
                        </div>
                        <div class="col-sm-3">
                            <div class="service-item">
                                <span class="fa-stack fa-4x hidden-xs">
                                <i class="fa fa-circle fa-stack-2x hidden-xs"></i>
                                <i class="fab fa-apple fa-stack-1x text-primary hidden-xs"></i>
                            </span>
                                <h4>
                                    <strong>macOS</strong>
                                </h4>
                            </div>
                        </div>
                    </div>
                    <!-- /.row (nested) -->
                </div>
                <!-- /.col-lg-10 -->
            </div>
            <!-- /.row -->
        </div>
        <!-- /.container -->
    </section>

    <!-- Callout 
    <aside class="callout">
        <div class="text-vertical-center">
            <h1>"Man, this Keith guy sure is a class act"</h1>
            <h3>--Anonymous</h3>
        </div>
    </aside>-->

    <!-- Portfolio -->
    <section id="portfolio" class="portfolio">
        <div class="container">
            <div class="row">
                <div class="col-lg-10 col-lg-offset-1 text-center">
                    <h2>My Projects</h2>
                    <hr class="small">
                    <div class="row">
                        <div class="col-md-6">
                            <div class="portfolio-item">
                                <a href="http://github.com/keithhand/personal-portal">
                                    <img class="img-portfolio img-responsive" src="img/portfolio-1.jpg">
                                </a>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="portfolio-item">
                                <a href="http://github.com/keithhand/ant-facts">
                                    <img class="img-portfolio img-responsive" src="img/portfolio-2.jpg">
                                </a>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="portfolio-item">
                                <a href="http://github.com/keithhand/se-project">
                                    <img class="img-portfolio img-responsive" src="img/portfolio-3.jpg">
                                </a>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="portfolio-item">
                                <a href="https://github.com/keithhand/scripts/blob/master/ucossh">
                                    <img class="img-portfolio img-responsive" src="img/portfolio-4.jpg">
                                </a>
                            </div>
                        </div>
                    </div>
                    <!-- /.row (nested) -->
                </div>
                <!-- /.col-lg-10 -->
            </div>
            <!-- /.row -->
        </div>
        <!-- /.container -->
    </section>

    <!-- Call to Action -->
    <div id="contact" class="call-to-action bg-primary">
        <div class="container">
            <div class="row">
                <div class="col-lg-12 text-center">
                    <h2>Contact Information</h2>
                </div>
            </div>
        </div>
    </div>

    <!-- Map -->
    <section class="map">
        <iframe width="100%" height="100%" frameborder="0" scrolling="no" marginheight="0" marginwidth="0" src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d494793.76199040335!2d-97.71640447924517!3d35.4437818739835!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x87ad8a547ef8d281%3A0x33a21274d14f3a9d!2sOklahoma+City%2C+OK!5e0!3m2!1sen!2sus!4v1555551250400!5m2!1sen!2sus"></iframe>
        </iframe>
    </section>

    <!-- Footer -->
    <footer>
        <div class="container">
            <div class="row">
                <div class="col-lg-10 col-lg-offset-1 text-center">
                    <h4><strong>Keith Hand</strong>
                    </h4>
                    <ul class="list-unstyled">
                        <li><i class="fa fa-phone fa-fw"></i> (405) 574-5813</li>
                        <li><i class="fa fa-envelope-o fa-fw"></i> <a href="mailto:rkeithhand@gmail.com">rkeithhand@gmail.com</a>
                        </li>
                    </ul>
                    <br>
                    <ul class="list-inline">
                        <li>
                            <a href="http://www.github.com/keithhand"><i class="fab fa-github fa-fw fa-3x"></i></a>
                        </li>
                    </ul>
                    <hr class="small">
                    <p class="text-muted">Copyright &copy; <?=date("Y")?> keithhand.xyz</p>
                </div>
            </div>
        </div>
        <a id="to-top" href="#top" class="btn btn-dark btn-lg"><i class="fa fa-chevron-up fa-fw fa-1x"></i></a>
    </footer>

    <!-- jQuery -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.0/jquery.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>

    <!-- Custom Theme JavaScript -->
    <script>
    // Closes the sidebar menu
    $("#menu-close").click(function(e) {
        e.preventDefault();
        $("#sidebar-wrapper").toggleClass("active");
    });
    // Opens the sidebar menu
    $("#menu-toggle").click(function(e) {
        e.preventDefault();
        $("#sidebar-wrapper").toggleClass("active");
    });




    // Scrolls to the selected menu item on the page
    // $(function() {
    //     $('a[href*=#]:not([href=#],[data-toggle],[data-target],[data-slide])').click(function() {
    //         if (location.pathname.replace(/^\//, '') == this.pathname.replace(/^\//, '') || location.hostname == this.hostname) {
    //             var target = $(this.hash);
    //             target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
    //             if (target.length) {
    //                 $('html,body').animate({
    //                     scrollTop: target.offset().top
    //                 }, 1000);
    //                 return false;
    //             }
    //         }
    //     });
    // });








    //#to-top button appears after scrolling
    var fixed = false;
    $(document).scroll(function() {
        if ($(this).scrollTop() > 250) {
            if (!fixed) {
                fixed = true;
                // $('#to-top').css({position:'fixed', display:'block'});
                $('#to-top').show("slow", function() {
                    $('#to-top').css({
                        position: 'fixed',
                        display: 'block'
                    });
                });
            }
        } else {
            if (fixed) {
                fixed = false;
                $('#to-top').hide("slow", function() {
                    $('#to-top').css({
                        display: 'none'
                    });
                });
            }
        }
    });
    // Disable Google Maps scrolling
    // See http://stackoverflow.com/a/25904582/1607849
    // Disable scroll zooming and bind back the click event
    var onMapMouseleaveHandler = function(event) {
        var that = $(this);
        that.on('click', onMapClickHandler);
        that.off('mouseleave', onMapMouseleaveHandler);
        that.find('iframe').css("pointer-events", "none");
    }
    var onMapClickHandler = function(event) {
            var that = $(this);
            // Disable the click handler until the user leaves the map area
            that.off('click', onMapClickHandler);
            // Enable scrolling zoom
            that.find('iframe').css("pointer-events", "auto");
            // Handle the mouse leave event
            that.on('mouseleave', onMapMouseleaveHandler);
        }
        // Enable map zooming with mouse scroll when the user clicks the map
    $('.map').on('click', onMapClickHandler);
    </script>

</body>

</html>
