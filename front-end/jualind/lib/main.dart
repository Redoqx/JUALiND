import 'package:flutter/material.dart';

import 'package:jualind/penjual/navigatorPenjual.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Admin",
      theme: ThemeData(
        primaryColor: Color.fromRGBO(139, 16, 255, 1),
      ),
      // home: LoginPage(),
      home: NavJual(),
    );
  }
}

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    var gradientBg = const LinearGradient(
      begin: Alignment.topCenter,
      end: Alignment.bottomCenter,
      colors: [
        Color.fromARGB(255, 115, 204, 255),
        Color.fromARGB(66, 127, 9, 206),
      ],
    );
    return Scaffold(
      body: SafeArea(
        child: Center(
          child: BgDanKotak(),
          // child: Container(
          //   decoration: BoxDecoration(
          //     gradient: gradientBg,
          //   ),
          //   child: Center(
          //     child: FractionallySizedBox(
          //       widthFactor: 0.8,
          //       heightFactor: 0.6,
          //     ),
          //   ),
          // ),
        ),
      ),
    );
  }
}

class BgDanKotak extends StatelessWidget {
  const BgDanKotak({super.key});

  @override
  Widget build(BuildContext context) {
    var gradientBg = const LinearGradient(
      begin: Alignment.topCenter,
      end: Alignment.bottomCenter,
      colors: [
        Color.fromARGB(255, 115, 204, 255),
        Color.fromARGB(66, 127, 9, 206),
      ],
    );
    return Container(
      decoration: BoxDecoration(
        gradient: gradientBg,
      ),
      child: Center(
        child: FractionallySizedBox(
          widthFactor: 0.8,
          heightFactor: 0.6,
          child: Container(
            padding: const EdgeInsets.all(20),
            decoration: const BoxDecoration(
              color: Colors.white,
              borderRadius: BorderRadius.all(Radius.circular(15)),
            ),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                Text(
                  "Masuk",
                  style: const TextStyle(
                    fontSize: 36,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                LoginForm()
              ],
            ),
          ),
        ),
      ),
    );
  }
}

class LoginForm extends StatefulWidget {
  const LoginForm({super.key});

  @override
  State<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  TextEditingController nameController = TextEditingController();
  TextEditingController passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: <Widget>[
            //username
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
              child: TextFormField(
                controller: nameController,
                decoration: const InputDecoration(
                  border: UnderlineInputBorder(),
                  labelText: 'Username',
                ),
              ),
            ),
            //password
            Container(
              margin: new EdgeInsets.only(bottom: 20),
              child: Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                child: TextFormField(
                  controller: passwordController,
                  decoration: const InputDecoration(
                    border: UnderlineInputBorder(),
                    labelText: 'Password',
                  ),
                ),
              ),
            ),
            //loginButton
            AnimatedContainer(
              duration: const Duration(milliseconds: 200),
              margin: new EdgeInsets.symmetric(vertical: 20.0),
              height: 60,
              curve: Curves.easeIn,
              decoration: BoxDecoration(
                color: Color.fromARGB(255, 197, 126, 255),
                borderRadius: BorderRadius.all(Radius.circular(60 / 2)),
              ),
              child: InkWell(
                onTap: () {
                  if (nameController.text == "aku" &&
                      passwordController.text == "passnya") {
                    print(nameController.text);
                    print(passwordController.text);
                  }
                },
                child: Center(
                    child: Text(
                  "MASUK",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    color: Colors.white,
                  ),
                )),
              ),
            ),
          ]),
    );
  }
}
