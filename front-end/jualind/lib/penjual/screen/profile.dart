import 'package:flutter/material.dart';

class ProfileJ extends StatefulWidget {
  const ProfileJ({super.key});

  @override
  State<ProfileJ> createState() => _ProfileJState();
}

class _ProfileJState extends State<ProfileJ> {
  // var isi =List<List<String>>[['d'],['c']];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: SingleChildScrollView(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: <Widget>[
          Container(),
        ],
      ),
    ));
  }
}
