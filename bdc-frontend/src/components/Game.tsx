import React, { Component } from 'react';
import Phaser from 'phaser';
import { IonPhaser } from '@ion-phaser/react';
import BoardPlugin from 'phaser3-rex-plugins/plugins/board-plugin.js';
//import { Board, HexagonGrid } from 'phaser3-rex-plugins/plugins/board-components.js';
//import KineticScrolling from 'phaser-kinetic-scrolling-plugin';
import KineticScroll from '../lib/KineticScroll';
const Random = Phaser.Math.Between;


class HexGridScene extends Phaser.Scene {

  board!: BoardPlugin.Board;
  print!: Phaser.GameObjects.Text;
  cameraController!: Phaser.Cameras.Controls.SmoothedKeyControl;
  rexBoard!: BoardPlugin;
  scrolling!: KineticScroll;
  origDragPoint?: Phaser.Math.Vector2;

  constructor() {
    super('hex-grid');
  }

  init() {
    this.cameras.main.setBackgroundColor('#24252A')
  }

  create() {
    var board = this.rexBoard.add.board({
      grid: {
        gridType: 'hexagonGrid',
        x: 50,
        y: 50,
        size: 50,
        staggeraxis: 'x',
        staggerindex: 'odd'
      },
      width: 20,
      height: 20
    }).forEachTileXY(function (tileXY, board) {
        var scene = board.scene as HexGridScene;
        var chess = scene.rexBoard.add.shape(board, tileXY.x, tileXY.y, 0, Random(0, 0xffffff), 0.7);
        scene.add.text(chess.x, chess.y, tileXY.x + ',' + tileXY.y)
            .setOrigin(0.5)
            .setTint(0x0);
    });

    board
        .setInteractive()
        .on('tiledown', function (pointer: any, tileXY: {x: number, y: number}) {
            console.log('down ' + tileXY.x + ',' + tileXY.y);
        })
        .on('tileup', function (pointer: any, tileXY: {x: number, y: number}) {
            console.log('up ' + tileXY.x + ',' + tileXY.y);
        })
        .on('tilemove', function (pointer: any, tileXY: {x: number, y: number}) {
            console.log('move ' + tileXY.x + ',' + tileXY.y);
        })
        .on('gameobjectdown', function (pointer: any, gameObject: any) {
            gameObject.setFillStyle(Random(0, 0xffffff), 0.7);
        })
        .on('tile1tap', function (tap: any, tileXY: {x: number, y: number}) {
            console.log('1 tap ' + tileXY.x + ',' + tileXY.y);
        })
        .on('tile2tap', function (tap: any, tileXY: {x: number, y: number}) {
            console.log('2 tap ' + tileXY.x + ',' + tileXY.y);
        })
        .on('tilepressstart', function (press: any, tileXY: {x: number, y: number}) {
            console.log('press start ' + tileXY.x + ',' + tileXY.y);
        })
        .on('tilepressend', function (press: any, tileXY: {x: number, y: number}) {
            console.log('press end ' + tileXY.x + ',' + tileXY.y);
        })
        .on('tileswipe', function (swipe: any, tileXY: {x: number, y: number}) {
            console.log(`swipe-${swipe.direction} ` + tileXY.x + ',' + tileXY.y);
        })

    this.board = board;
    this.print = this.add.text(10, 10, '').setScrollFactor(0);


    /*var cursors = this.input.keyboard.createCursorKeys();
    this.cameraController = new Phaser.Cameras.Controls.SmoothedKeyControl({
        camera: this.cameras.main,

        left: cursors.left,
        right: cursors.right,
        up: cursors.up,
        down: cursors.down,
        zoomIn: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.A),
        zoomOut: this.input.keyboard.addKey(Phaser.Input.Keyboard.KeyCodes.E),

        acceleration: 0.06,
        drag: 0.003,
        maxSpeed: 0.3
    });*/

    this.scrolling.configure({
      kineticMovement: true,
      horizontalScroll: true,
      verticalScroll: true
    });
  }

  update(time: any, delta: any) {
    if (this.game.input.activePointer.isDown) {
      if (this.origDragPoint) {
        this.scrolling.move(this.origDragPoint, this.game.input.activePointer.position.x, this.game.input.activePointer.position.y);
      }
      this.origDragPoint = this.game.input.activePointer.position.clone();
      this.scrolling.beginMove(this.origDragPoint);
    } else {
      this.origDragPoint = undefined;
      this.scrolling.endMove();
    }
    /*if (this.game.input.activePointer.isDown) {
      if (this.origDragPoint) {
        this.cameras.main.x += this.origDragPoint.x - this.game.input.activePointer.position.x;
        this.cameras.main.y += this.origDragPoint.y - this.game.input.activePointer.position.y;
      }
      this.origDragPoint = this.game.input.activePointer.position.clone();
    } else {
      this.origDragPoint = undefined;
    }

    this.cameraController.update(delta);
    */
    var pointer = this.input.activePointer;
    var out = this.board.worldXYToTileXY(pointer.worldX, pointer.worldY, true);
    this.print.setText(out.x + ',' + out.y);
  }
}

class Game extends Component {

  config: Phaser.Types.Core.GameConfig;

  constructor() {
    super({});
    
    this.config = {
      width: '100%',
      height: '100%',
      type: Phaser.AUTO,
      scene: HexGridScene,
      plugins: {
        scene: [{
          key: 'rexBoard',
          plugin: BoardPlugin,
          mapping: 'rexBoard'
        },{
          key: 'scrolling',
          plugin: KineticScroll,
          mapping: 'scrolling'
        }]
      }
    }
  }
  
  render() {
    return (
      <IonPhaser game={this.config} initialize={true} />
    );
  }
}

export default Game;
