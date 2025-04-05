import { ref } from 'vue'
import axios from 'axios'

export interface Coordinate {
  latitude: number
  longitude: number
}

export interface RouteInfo {
  distance: number
  duration: number
  coordinates: Coordinate[]
}

export interface Weather {
  temperature: number
  feelsLike: number
  condition: string
  windSpeed: number
  humidity: number
}

export interface MapSettings {
  center: Coordinate
  zoom: number
  mapType: 'map' | 'satellite' | 'hybrid'
  showTraffic: boolean
}

export interface TrackPoint {
  coordinate: Coordinate
  timestamp: number
  elevation: number
  speed: number
  accuracy: number
}

class MapService {
  private map: any = null
  private route: any = null
  private trackPoints: TrackPoint[] = []
  private isTracking = ref(false)
  private watchId: number | null = null

  constructor(private apiKey: string) {}

  async initMap(containerId: string, settings: MapSettings): Promise<void> {
    // @ts-ignore
    await ymaps.ready()
    
    // @ts-ignore
    this.map = new ymaps.Map(containerId, {
      center: [settings.center.latitude, settings.center.longitude],
      zoom: settings.zoom,
      type: settings.mapType,
      controls: ['zoomControl', 'fullscreenControl']
    })

    if (settings.showTraffic) {
      // @ts-ignore
      this.map.controls.add('trafficControl')
    }
  }

  async getRouteInfo(coordinates: Coordinate[]): Promise<RouteInfo> {
    const response = await axios.post('/api/maps/route', coordinates)
    return response.data
  }

  async getElevation(coordinates: Coordinate[]): Promise<number[]> {
    const response = await axios.post('/api/maps/elevation', coordinates)
    return response.data
  }

  async getWeather(coordinate: Coordinate): Promise<Weather> {
    const response = await axios.get('/api/maps/weather', { params: coordinate })
    return response.data
  }

  startTracking(): void {
    if (this.isTracking.value) return

    this.isTracking.value = true
    this.trackPoints = []

    if (navigator.geolocation) {
      this.watchId = navigator.geolocation.watchPosition(
        (position) => {
          const trackPoint: TrackPoint = {
            coordinate: {
              latitude: position.coords.latitude,
              longitude: position.coords.longitude
            },
            timestamp: Date.now(),
            elevation: position.coords.altitude || 0,
            speed: position.coords.speed || 0,
            accuracy: position.coords.accuracy
          }

          this.trackPoints.push(trackPoint)
          this.updateRoute()
        },
        (error) => {
          console.error('Error getting location:', error)
        },
        {
          enableHighAccuracy: true,
          timeout: 5000,
          maximumAge: 0
        }
      )
    }
  }

  stopTracking(): void {
    if (!this.isTracking.value) return

    this.isTracking.value = false
    if (this.watchId !== null) {
      navigator.geolocation.clearWatch(this.watchId)
      this.watchId = null
    }
  }

  private updateRoute(): void {
    if (this.trackPoints.length < 2) return

    const coordinates = this.trackPoints.map(point => [
      point.coordinate.latitude,
      point.coordinate.longitude
    ])

    if (this.route) {
      this.map.geoObjects.remove(this.route)
    }

    // @ts-ignore
    this.route = new ymaps.Polyline(coordinates, {
      balloonContent: 'Маршрут тренировки'
    }, {
      strokeColor: '#FF0000',
      strokeWidth: 4,
      strokeOpacity: 0.8
    })

    this.map.geoObjects.add(this.route)
    this.map.setBounds(this.route.geometry.getBounds())
  }

  getTrackPoints(): TrackPoint[] {
    return this.trackPoints
  }

  clearTrack(): void {
    this.trackPoints = []
    if (this.route) {
      this.map.geoObjects.remove(this.route)
      this.route = null
    }
  }
}

export const mapService = new MapService(import.meta.env.VITE_YANDEX_MAPS_API_KEY) 