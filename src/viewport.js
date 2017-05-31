//@flow

import { View, MouseEvent, Point, Layer } from 'paper';


/**
 * Encapsulate the viewport's logic.
 * Expose methods to easily define how the current floorplan should be shown
 * to the user.
 */
export default class Viewport {

  view: View;

  _mainLayer: Layer;
  _rasterLayer: Layer;
  _shouldRaster: bool;

  constructor(view: View) {
    this.view = view;
    this.view.autoUpdate = false;

    this._mainLayer = view._project.activeLayer;
    this._rasterLayer = view._project.addLayer(new Layer());
    this._shouldRaster = false;

    this.view.onMouseDrag = (e) => this._translate(e);
    this.view.onMouseUp = (e) => {
      this._shouldRaster = false;
      this.redraw();
    }
  }

  /**
   * Set the viewport's zoom factor.
   */
  setZoom(value: number) {
    this.view.setZoom(value);
  }

  /**
   * Set the viewport's center at provided position.
   */
  setCenter(x: number, y: number) {
    this.view.center = new Point(x, y);
  }

  /**
   * redraw the viewport's display.
   */
  redraw() {
    requestAnimationFrame(() => {
      if (this._shouldRaster) {
        if(this._rasterLayer.isEmpty())
          this._rasterLayer.addChild(this._mainLayer.rasterize());

        // hide the main layer and activate the raster layer.
        this._mainLayer.visible = false;
        this._rasterLayer.visible = true;
      } else {
        // hide the raster layer and activate the main layer.
        this._rasterLayer.visible = false;
        this._mainLayer.visible = true;
      }

      this.view.requestUpdate(this.view.update())
    })
  }

  /**
   * Implement how the viewport should translate.
   */
  _translate(e: MouseEvent) {
    const delta = e.delta;
    delta.length /= 2;  // reduce velocity of the drag

    this.view.center = this.view.center.add(new Point(-delta.x, -delta.y));

    // Raster the main layer when translating the viewport
    this._shouldRaster = true;

    // redraw the viewport
    this.redraw();
  }
}
