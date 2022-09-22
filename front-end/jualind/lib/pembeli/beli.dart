import 'package:flutter/material.dart';
// import 'package:flutter_staggered_grid_view/flutter_staggered_grid_view.dart';
// import 'package:scaffold_gradient_background/scaffold_gradient_background.dart';

class beli extends StatefulWidget {
  const beli({super.key});

  @override
  State<beli> createState() => _beliState();
}

class _beliState extends State<beli> {
  TextEditingController cariin = TextEditingController();

  bool tampilNama = true;
  @override
  Widget build(BuildContext context) {
    var gradientBg = const LinearGradient(
      begin: Alignment.topRight,
      end: Alignment.bottomLeft,
      colors: [
        Color.fromARGB(255, 115, 204, 255),
        Color.fromARGB(66, 127, 9, 206),
      ],
    );

    return Scaffold(
      body: Container(
        decoration: BoxDecoration(
          gradient: gradientBg,
        ),
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(16),
              child: Row(
                children: [
                  Expanded(
                    child: Text(
                      "JUALiND",
                      style: TextStyle(
                        fontSize: 32,
                        fontWeight: FontWeight.w900,
                      ),
                    ),
                  ),
                  IconButton(
                    onPressed: () {},
                    icon: Icon(Icons.search),
                  ),
                ],
              ),
            ),
            Expanded(
              child: GridView.count(crossAxisCount: 2),
            ),
          ],
        ),

        // child: Column(
        //   mainAxisSize: MainAxisSize.min,
        //   children: <Widget>[
        //     Padding(
        //       padding: EdgeInsets.all(10),
        //       child: Container(
        //         child: Row(
        //           mainAxisAlignment: MainAxisAlignment.start,
        //           crossAxisAlignment: CrossAxisAlignment.stretch,
        //           children: <Widget>[
        //             CircleAvatar(
        //               backgroundColor: Colors.white,
        //               radius: 10,
        //               backgroundImage: null,
        //             ),
        //             Text(
        //               tampilNama ? "Halo, Siapa ya?" : '',
        //               style: TextStyle(fontSize: 5),
        //             ),
        //             AnimatedContainer(
        //               duration: Duration(milliseconds: 500),
        //               child: tampilNama
        //                   ? Icon(
        //                       Icons.search,
        //                     )
        //                   : TextField(
        //                       controller: cariin,
        //                       onSubmitted: (cariin) {
        //                         //panggil API pencarian
        //                       },
        //                     ),
        //             )
        //           ],
        //         ),
        //       ),
        //     ),
        //   ],
        // ),
      ),
    );
  }
}
