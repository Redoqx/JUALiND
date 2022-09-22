import 'package:flutter/material.dart';

import 'package:jualind/penjual/screen/dasboard.dart';
import 'package:jualind/penjual/screen/Jualin.dart';
import 'package:jualind/penjual/screen/profile.dart';

class NavJual extends StatefulWidget {
  const NavJual({super.key});

  @override
  State<NavJual> createState() => _NavJualState();
}

class _NavJualState extends State<NavJual> {
  int _selectedIndex = 0;
  static const List<Widget> _widgetOptions = <Widget>[
    HomeJ(),
    Jualin(),
    ProfileJ(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: _widgetOptions.elementAt(_selectedIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.bar_chart),
            label: 'Dashboard',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.add),
            label: 'Jualind',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
        currentIndex: _selectedIndex,
        backgroundColor: Color(0xFF8B10FF),
        selectedItemColor: Color.fromARGB(255, 234, 255, 236),
        onTap: _onItemTapped,
      ),
    );
  }
}
