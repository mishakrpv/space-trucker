import { TestBed } from '@angular/core/testing';

import { SecretsManagerService } from './secrets-manager.service';

describe('SecretsManagerService', () => {
  let service: SecretsManagerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SecretsManagerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
