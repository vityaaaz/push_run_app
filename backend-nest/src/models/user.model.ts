import { ApiProperty } from '@nestjs/swagger';

export class User {
  @ApiProperty()
  id: string;

  @ApiProperty()
  email: string;

  @ApiProperty()
  firstName: string;

  @ApiProperty()
  lastName: string;

  @ApiProperty()
  username: string;

  @ApiProperty()
  avatar: string;

  @ApiProperty()
  bio: string;

  @ApiProperty()
  totalDistance: number;

  @ApiProperty()
  totalTime: number;

  @ApiProperty()
  level: number;

  @ApiProperty()
  xp: number;

  @ApiProperty()
  createdAt: Date;

  @ApiProperty()
  updatedAt: Date;

  @ApiProperty({ required: false })
  githubId?: number;

  @ApiProperty({ required: false })
  githubLogin?: string;

  @ApiProperty({ required: false })
  accessToken?: string;
}

export class RegisterRequest {
  @ApiProperty()
  email: string;

  @ApiProperty()
  password: string;

  @ApiProperty()
  firstName: string;

  @ApiProperty()
  lastName: string;

  @ApiProperty()
  username: string;
}

export class LoginRequest {
  @ApiProperty()
  email: string;

  @ApiProperty()
  password: string;
}

export class UpdateProfileRequest {
  @ApiProperty({ required: false })
  firstName?: string;

  @ApiProperty({ required: false })
  lastName?: string;

  @ApiProperty({ required: false })
  username?: string;

  @ApiProperty({ required: false })
  avatar?: string;

  @ApiProperty({ required: false })
  bio?: string;
}

export class GitHubUser {
  @ApiProperty()
  id: number;

  @ApiProperty()
  login: string;

  @ApiProperty()
  email: string;

  @ApiProperty()
  name: string;

  @ApiProperty()
  avatar_url: string;
} 